const {InjectManifest} = require('workbox-webpack-plugin');

const SIZE_5MB = 5 * 1024 * 1024;

module.exports = function override(config, env) {
    config.plugins.push(
        new InjectManifest({
            swSrc: './src/sw/firebase-messaging-sw.ts',
            swDest: 'firebase-messaging-sw.js',
            maximumFileSizeToCacheInBytes: SIZE_5MB,
        })
    );
    config.plugins.push(
        new InjectManifest({
            swSrc: './src/sw/cache-sw.ts',
            swDest: 'cache-sw.js',
            maximumFileSizeToCacheInBytes: SIZE_5MB,
        })
    );
    return config;
};