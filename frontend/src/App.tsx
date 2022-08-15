import React, {useEffect, useState} from 'react';
import "@fontsource/roboto-mono";
import {Home} from "./Page/Home";
import {Box, Container, ThemeProvider} from "@mui/system";
import {CssBaseline, Snackbar} from '@mui/material';
import createTheme, {ThemeMode} from "./theme";
import WbIncandescentOutlinedIcon from '@mui/icons-material/WbIncandescentOutlined';
import Version from "./Component/Version";
import ApiClient from "./Api/ApiClient";

type OnlineStatus = false | true | null;

function App() {

    let defaultThemeMode: ThemeMode = "light";
    if (window.matchMedia && window.matchMedia('(prefers-color-scheme: dark)').matches) {
        defaultThemeMode = "dark"
    }
    const savedThemeMode = localStorage.getItem("themeMode");
    if (savedThemeMode != null && (savedThemeMode === 'dark' || savedThemeMode === 'light')) {
        defaultThemeMode = savedThemeMode;
    }

    const [themeMode, setThemeMode] = useState<ThemeMode>(defaultThemeMode)

    const theme = createTheme(themeMode);

    const toggleMode = () => {
        const mode: ThemeMode = themeMode === 'light' ? 'dark' : 'light';
        localStorage.setItem("themeMode", mode)
        setThemeMode(mode);
    }

    const [srvConn, setSrvConn] = useState<OnlineStatus>(null);
    const [openSnackbar, setOpenSnackbar] = useState<boolean>(false);

    useEffect(() => {
        if (srvConn === null) {
            ApiClient.get('/health').then(data => {
                if (data.status === 200) {
                    setSrvConn(true);
                } else {
                    setSrvConn(false);
                    setOpenSnackbar(true);
                }
            }).catch(err =>{
                console.error(err)
                setSrvConn(false);
                setOpenSnackbar(true);
            })
        }
    }, [srvConn]);

    return (
        <ThemeProvider theme={theme}>
            <Snackbar
                anchorOrigin={{ vertical: "bottom", horizontal: "center" }}
                key={`notConnect`}
                open={openSnackbar}
                autoHideDuration={3000}
                message="Connection lost"
                onClose={() => {setOpenSnackbar(false)}}
            />
            <CssBaseline />
            <Container
                maxWidth="md"
                component="div"
                sx={{
                    minHeight: '100vh',
                    display: 'grid',
                    gridTemplateRows: 'auto 1fr auto',
                }}
            >
                <Box sx={{
                    pt: 2,
                    display: 'flex',
                    justifyContent:'right',
                }}>

                    <WbIncandescentOutlinedIcon onClick={toggleMode} sx={{
                        fontSize: '2.6em',
                        cursor: 'pointer',
                    }}/>
                </Box>
                <Home />
                <Box sx={{
                    minHeight: 32,
                }}>
                    <Version />
                </Box>

            </Container>
        </ThemeProvider>
    );
}

export default App;
