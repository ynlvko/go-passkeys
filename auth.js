function base64URLStringToArrayBuffer(base64URLString) {
    let padding = '='.repeat((4 - base64URLString.length % 4) % 4);
    let base64 = (base64URLString + padding)
        .replace(/-/g, '+')
        .replace(/_/g, '/');

    let rawData = window.atob(base64);
    let outputArray = new Uint8Array(rawData.length);

    for (let i = 0; i < rawData.length; ++i) {
        outputArray[i] = rawData.charCodeAt(i);
    }
    return outputArray.buffer;
}

function arrayBufferToBase64String(arrayBuffer) {
    let byteArray = new Uint8Array(arrayBuffer);
    let byteString = '';
    for (let i = 0; i < byteArray.byteLength; i++) {
        byteString += String.fromCharCode(byteArray[i]);
    }
    return window.btoa(byteString).replace(/\+/g, '-').replace(/\//g, '_').replace(/=/g, '');
}

// Begin registration
async function beginRegistration() {
    const response = await fetch('https://go-passkeys.onrender.com/begin-registration');
    const data = await response.json();

    // Convert challenge and user ID from Base64 to ArrayBuffer
    data.publicKey.challenge = base64URLStringToArrayBuffer(data.publicKey.challenge);
    if (data.publicKey.user) {
        data.publicKey.user.id = base64URLStringToArrayBuffer(data.publicKey.user.id);
    }

    return data;
}

// Finish registration
async function finishRegistration(creds) {
    // Convert ArrayBuffer to Base64
    const authData = creds.response.attestationObject;
    const clientDataJSON = creds.response.clientDataJSON;
    const registrationData = {
        id: creds.id,
        rawId: arrayBufferToBase64String(creds.rawId),
        type: creds.type,
        response: {
            attestationObject: arrayBufferToBase64String(authData),
            clientDataJSON: arrayBufferToBase64String(clientDataJSON),
        },
    };

    await fetch('https://go-passkeys.onrender.com/finish-registration', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(registrationData),
    });
}

// Run registration flow
async function runRegistrationFlow() {
    try {
        const options = await beginRegistration();
        const creds = await navigator.credentials.create(options);
        if (creds) {
            await finishRegistration(creds);
            console.log('Registration successful');
        } else {
            console.log('Registration failed');
        }
    } catch (error) {
        console.error('Error:', error);
    }
}

runRegistrationFlow();
