import ArrowForwardIosIcon from "@mui/icons-material/ArrowForwardIos";
import { Box, Container, Divider, Grid, Paper, Stack, Typography } from "@mui/material";
import React from "react";
import Link from "@mui/material/Link";
import { Copyright } from "../Common/Copyright";

const Footer = () => {
  return (
    <Paper
      sx={{ marginTop: "calc(10% + 60px)", bottom: 0,marginBottom: 0}}
      component="footer"
      square
      variant="outlined"
    >
      <Box bgcolor={"#1976d2"}>
        <Container
          maxWidth="lg"
          sx={{
            py: 5,
          }}
        >
          <Grid container spacing={5}>
            <Grid item xs={12} sm={4}>
              <Typography variant="h5" component={"h5"} sx={{ mb: 2 }}>
                About us
              </Typography>
              <Typography variant="body1" component={"p"} maxWidth="15rem">
                Bookhub is a platform where you can find your favorite books and
                authors. You can also add your own books and authors to the
                platform.
              </Typography>
            </Grid>
            <Grid item xs={12} sm={4}>
              <Typography variant="h5" component={"h5"} sx={{ mb: 2 }}>
                Useful Links
              </Typography>
              <Stack direction="row" alignItems="center" marginBottom={"1rem"}>
                <ArrowForwardIosIcon />
                <Typography variant="h6">
                  <Link href="/" underline="none" color={"black"}>
                    {"Home"}
                  </Link>
                </Typography>
              </Stack>
              <Stack direction="row" alignItems="center" marginBottom={"1rem"}>
                <ArrowForwardIosIcon />
                <Typography variant="h6">
                  <Link href="/login" underline="none" color={"black"}>
                    {"Log In"}
                  </Link>
                </Typography>
              </Stack>
              <Stack direction="row" alignItems="center" marginBottom={"1rem"}>
                <ArrowForwardIosIcon />
                <Typography variant="h6">
                  <Link href="/signup" underline="none" color={"black"}>
                    {"Sign Up"}
                  </Link>
                </Typography>
              </Stack>
            </Grid>
            <Grid item xs={12} sm={4}>
              <Typography variant="h5" component={"h5"} sx={{ mb: 2 }}>
                Contact Us
              </Typography>
              <Typography variant="body1" component={"p"} maxWidth="5rem">
                Kharghar Sector 35 Maharashtra,410210 India
              </Typography>
            </Grid>
          </Grid>
          <Divider/>
          <Copyright sx={{mt:4}}></Copyright>
        </Container>
      </Box>
    </Paper>
  );
};

export default Footer;
