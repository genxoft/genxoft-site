export const getMessagingToken = async (setTokenFound: (token: string | null) => void) => {
    console.error('You need configure fcm');
    setTokenFound(null);
    return;
}

export const firebaseApp = null;