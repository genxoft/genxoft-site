import {Box} from "@mui/material";
import {styled} from "@mui/system";

export default function Version() {

    const version = process.env.REACT_APP_VERSION == undefined ? 'dev' : process.env.REACT_APP_VERSION;
    const build = process.env.REACT_APP_BUILD == undefined ? 'dev' : process.env.REACT_APP_BUILD;
    const build_mode = process.env.REACT_APP_BUILD_MODE == undefined ? 'dev' : process.env.REACT_APP_BUILD_MODE;

    const VersionSign = styled('span')(
        () => `
  opacity: 0.2;
`,
    );

    return (
        <Box sx={{
            display: 'flex',
            justifyContent: 'right',
        }}>
            <VersionSign>{version}-{build_mode} Build: {build}</VersionSign>
        </Box>
    )
}