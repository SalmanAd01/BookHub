import { Typography } from "@mui/material";

export function Copyright(props) {
    console.log(props);
    return (
      <Typography variant="body2" color="text.secondary" align="center" {...props}>
        {'Copyright © '}
        <a color="inherit" href="https://bookkhub.herokuapp.com/">
          Bookhub
        </a>{' '}
        {new Date().getFullYear()}
        {'.'}
      </Typography>
    );
  }