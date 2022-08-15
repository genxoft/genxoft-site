import {precacheAndRoute} from 'workbox-precaching';
import {clientsClaim} from "workbox-core";

declare const self: ServiceWorkerGlobalScope;

precacheAndRoute(self.__WB_MANIFEST);

clientsClaim();

self.addEventListener('message', (event) => {
    const notificationOptions = {
        body: 'Background Message body.',
        icon: '/img/logo256.png'
    };
    self.registration.showNotification("Message received", notificationOptions);
});