import {Box, Container, styled} from "@mui/system";
import React, {useState} from "react";
import {getMessagingToken} from "../config/firebaseInit";
import TelegramIcon from '@mui/icons-material/Telegram';
import TwitterIcon from '@mui/icons-material/Twitter';
import GitHubIcon from '@mui/icons-material/GitHub';
import ApiClient from "../Api/ApiClient";

const Skill = styled('span')(
    ({ theme }) => `
  background-color: ${theme.palette.primary.light};
  border-radius: ${theme.shape.borderRadius}px;
  padding:0 ${theme.spacing(1)};
  margin:0;
  color: ${theme.palette.primary.contrastText};
`);

const InlineButton = styled('button')(
    ({ theme}) => `
    color: ${theme.palette.success.main};
    cursor: pointer;
    text-decoration: underline;
    font-size:1em;
    background: none;
    padding:0;
    border:none;
    margin:0;
`);

const ContactBlock = styled('span')(
    () => `
    display: inline-block;
    white-space: nowrap;
`);

export function Home() {

    const [showSoftSkills, setShowSoftSkills] = useState<boolean>(false);

    const display = showSoftSkills ? 'inline' : 'none';

    const SkillSoft = styled('span')(
        ({ theme }) => `
  display: ${display};
  background-color: ${theme.palette.secondary.light};
  border-radius: ${theme.shape.borderRadius}px;
  padding:0 ${theme.spacing(1)};
  margin:0;
  color: ${theme.palette.secondary.contrastText};
  opacity: 0.6;
`);

    const onClickMore = () => {
        setShowSoftSkills(true);
    }

    const onClickLess = () => {
        setShowSoftSkills(false);
    }

    const subscribeNotifications = () => {
        getMessagingToken((token) => {
            const timezone = Intl.DateTimeFormat().resolvedOptions().timeZone;
            console.log(token);
            ApiClient.post('/fcm-token', {
                token,
                timezone: timezone,
            }).then((data) => {
                if (data.status !== 200) {
                    console.error("Server returns" + data.status + data.statusText);
                }
            }).catch((err) => {
                console.error(err);
            })
        }).then();
    }

    return (

        <Container
            maxWidth="md"
            component="div"

            sx={{
                pt: 8,
                fontSize: {xs: "12pt", md: "16pt"}
            }}
        >
            <Box>
                <p>🧑‍💻 Backend developer</p>
                <h1>👋&nbsp;Hello. I&nbsp;am&nbsp;Simon</h1>
                <p>
                    If you want to contact_me you can do it with{" "}
                    <ContactBlock>
                        <TelegramIcon sx={{verticalAlign: 'middle', mr: 0.3}} />
                        <a href="https://t.me/genxoft">genxoft</a>
                    </ContactBlock>{" "}
                    or{" "}
                    <ContactBlock>
                        <TwitterIcon sx={{verticalAlign: 'middle', mr: 0.3}} />
                        <a href="https://twitter.com/genx_ru">genx_ru</a>
                    </ContactBlock>{" "}<br />
                    and also{" "}
                    <ContactBlock>
                        <GitHubIcon sx={{verticalAlign: 'middle', mr: 0.3}} /><a href="https://github.com/genxoft">genxoft</a>
                    </ ContactBlock>.
                </p>
                <Box component="p" sx={{
                    lineHeight: 1.8,
                }}>
                    My skills:{" "}
                    <Skill>Go</Skill>{" "}<Skill>PHP</Skill>{" "}<Skill>REST</Skill>{" "}<Skill>SOA</Skill>{" "}
                    <Skill>SQL</Skill>{" "}<Skill>NoSQL</Skill>{" "}<SkillSoft>Yii2</SkillSoft>{" "}
                    <SkillSoft>MySQL</SkillSoft>{" "}<SkillSoft>Redis</SkillSoft>{" "}<SkillSoft>DoctrineORM</SkillSoft>{" "}
                    <SkillSoft>Git</SkillSoft>{" "}<SkillSoft>Docker</SkillSoft>{" "}<SkillSoft>RabbitMQ</SkillSoft>{" "}
                    <SkillSoft>PostgreSQL</SkillSoft>{" "}<SkillSoft>Swagger</SkillSoft>{" "}<SkillSoft>React</SkillSoft>{" "}
                    <SkillSoft>TypeScrypt</SkillSoft>{" "}<SkillSoft>HTML5</SkillSoft>{" "}<SkillSoft>CSS</SkillSoft>{" "}<SkillSoft>MaterialUI</SkillSoft>{" "}
                    <InlineButton onClick={onClickMore} sx={{display: showSoftSkills ? 'none' : 'inline'}}>more...</InlineButton>
                    <InlineButton onClick={onClickLess} sx={{display: showSoftSkills ? 'inline' : 'none'}}>less...</InlineButton>
                </Box>
                <p>About me: 💻🛵🧳🕹️</p>
                <p>If you want to receive weird notifications from time to time&nbsp;<InlineButton onClick={subscribeNotifications}>click&nbsp;here</InlineButton>.</p>
            </Box>
        </Container>

    );
}
