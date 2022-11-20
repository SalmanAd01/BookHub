import React from "react";
import Image from "mui-image";
import { Box, Button, Container, createTheme, Typography } from "@mui/material";
const Hero = () => {
  const theme = createTheme({
    breakpoints: {
      values: {
        xs: 0,
        sm: 600,
        md: 960,
        lg: 1280,
        xl: 1920,
      },
    },
  });
  theme.typography = {
    h1: {
      fontSize: "3rem",
      [theme.breakpoints.up("sm")]: {
        fontSize: "4rem",
      },
      [theme.breakpoints.up("md")]: {
        fontSize: "5rem",
      },
      [theme.breakpoints.up("lg")]: {
        fontSize: "6rem",
      },
    },
    h2: {
      fontSize: "2rem",
      [theme.breakpoints.up("sm")]: {
        fontSize: "2.5rem",
      },
      [theme.breakpoints.up("md")]: {
        fontSize: "3.5rem",
      },
      [theme.breakpoints.up("lg")]: {
        fontSize: "3rem",
      },
    },
    p: {
      fontSize: "1rem",
      [theme.breakpoints.up("sm")]: {
        fontSize: "1.5rem",
      },
      [theme.breakpoints.up("md")]: {
        fontSize: "2rem",
      },
      [theme.breakpoints.up("lg")]: {
        fontSize: "2rem",
      },
    },
  };
  return (
    <Container>
      <Box
        sx={{
          display: "flex",
          flexDirection:{ xs: "column-reverse", md: "row"},
          justifyContent: "center",
          alignItems: "center",
          height: "75vh",
          marginTop : {xs:'2rem'}
        }}
      >
        <Box theme={theme}>
          <Typography variant="h2" component="h2" theme={theme}>
            Get University Books For Free BookHub
          </Typography>
          <Typography variant="p" component="p" theme={theme}>
            A Place To Share And Download The Books
          </Typography>
          <Button
            variant="contained"
            sx={{
              backgroundColor: "#3f51b5",
              color: "#fff",
              marginTop: "1rem",
            }}
          >
            Get Started
          </Button>
        </Box>
        <Box theme={theme}>
          <Image
            src="https://bookkhub.herokuapp.com/assets/img/resize-1641925062449711649ManReadingaBookVectorIllustration1removebgpreview-removebg-preview.png"
            showLoading={true}
            height="50vh"
          />
        </Box>
      </Box>
    </Container>
  );
};

export default Hero;
