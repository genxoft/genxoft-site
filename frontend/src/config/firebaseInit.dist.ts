export const getMessagingToken = async (setTokenFound: (token: string | null) => void) => {
    console.log('You need configure fcm');
    setTokenFound(null);
    return;
}

export const firebaseApp = null;