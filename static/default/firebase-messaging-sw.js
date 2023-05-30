// Give the service worker access to Firebase Messaging.
// Note that you can only use Firebase Messaging here. Other Firebase libraries
// are not available in the service worker.
importScripts('https://www.gstatic.com/firebasejs/8.10.0/firebase-app.js');
importScripts('https://www.gstatic.com/firebasejs/8.10.0/firebase-messaging.js');

// missing databaseURL: 'https://project-id.firebaseio.com',
const firebaseConfig = {
    apiKey: "AIzaSyBf7AwqeQMMUmzAm8uRZaWidhiif1iuZ-w",
    authDomain: "mystudy-f4eaa.firebaseapp.com",
    projectId: "mystudy-f4eaa",
    storageBucket: "mystudy-f4eaa.appspot.com",
    messagingSenderId: "898225550263",
    appId: "1:898225550263:web:41fb5dc4f38afea8572e32",
    measurementId: "G-8FN1H0KLF7"
};
// Initialize the Firebase app in the service worker by passing in
// your app's Firebase config object.
// https://firebase.google.com/docs/web/setup#config-object
firebase.initializeApp(firebaseConfig);

// Retrieve an instance of Firebase Messaging so that it can handle background
// messages.
const messaging = firebase.messaging();

messaging.onBackgroundMessage((payload) => {
    debugger
    console.log('[firebase-messaging-sw.js] Received background message ', payload);
    // Customize notification here
    const notificationTitle = 'Background Message Title';
    const notificationOptions = {
      body: 'Background Message body.',
      icon: '/firebase-logo.png'
    };
  
    self.registration.showNotification(notificationTitle,
      notificationOptions);
  });