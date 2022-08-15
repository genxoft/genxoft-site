import { createTheme as createThemeSuper} from '@mui/material';
import * as React from 'react';
import darkScrollbar from '@mui/material/darkScrollbar';
export type ThemeMode = 'dark' | 'light'

export const ThemeContext = React.createContext<{themeMode: ThemeMode, setThemeMode: (() => {}) | undefined}>({
    themeMode: 'dark',
    setThemeMode: undefined
})

const createTheme = (mode: ThemeMode) => createThemeSuper({

    palette: {
        mode: mode,
        background: {
            default: mode === 'dark' ? '#000f1a' : '#ffffff',
            paper: '#cccccc'
        },
        primary: {
            main: '#ff7f00',
            light: '#ff9d3f',
            dark: '#b53d00',
            contrastText: '#000000'
        },
        secondary: {
            main: '#455a64',
            light: '#718792',
            dark: '#1c313a',
            contrastText: '#ffffff'
        },
        success: {
            main: '#c62828',
            light: '#ff5f52',
            dark: '#8e0000',
            contrastText: '#ffffff'
        }
    },
    shape: {
        borderRadius: 8
    },
    components: {
        MuiCssBaseline: {
            styleOverrides: (themeParam) => ({
                body: {
                    ...(themeParam.palette.mode === 'dark' ? darkScrollbar() : null),
                    background: themeParam.palette.background.default,
                },
                a: {
                    color: themeParam.palette.success.main
                }
            })
        }
    }

});

export default createTheme;
