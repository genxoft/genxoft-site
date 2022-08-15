import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import App from './App';
import {registerCacheSw} from "./sw/Registration";
import {firebaseApp} from "./config/firebaseInit";
import {getMessaging, onMessage} from "firebase/messaging";

const root = ReactDOM.createRoot(
  document.getElementById('root') as HTMLElement
);
root.render(
  <React.StrictMode>
    <App />
  </React.StrictMode>
);

registerCacheSw();
//registerMessagesSw();

if (firebaseApp !== null) {
    console.log(firebaseApp);
    const messaging = getMessaging(firebaseApp);
    console.log(messaging);
    onMessage(messaging, (payload) => {
        console.log('Message received. ', payload);
        // TODO: show message
    });
}
