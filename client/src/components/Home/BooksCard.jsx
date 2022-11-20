import * as React from "react";
import Card from "@mui/material/Card";
import CardActions from "@mui/material/CardActions";
import CardContent from "@mui/material/CardContent";
import CardMedia from "@mui/material/CardMedia";
import Button from "@mui/material/Button";
import Typography from "@mui/material/Typography";
import { Container, Grid } from "@mui/material";
import { Download } from "@mui/icons-material";

export default function BooksCard() {
  const Cards = [];
  for (let i = 0; i < 9; i++) {
    Cards.push(
      <Card key={i} sx={{ maxWidth: 345 }}>
        <CardMedia
          component="img"
          height="140"
          image="https://www.bookwalas.com/wp-content/uploads/2021/02/41aYxb6k9XL._SX351_BO1204203200_.jpg"
          alt="green iguana"
          sx={{ objectFit: "contain" }}
        />
        <CardContent>
          <Typography gutterBottom variant="h5" component="div">
            Mathematics-Sem1
          </Typography>
          <Typography variant="body2" color="text.secondary">
            Computer Engineer-Mumbai University
          </Typography>
          <Typography variant="p" color="text.secondary">
            GV Kumbhojkar
          </Typography>
        </CardContent>
        <CardActions>
          <Button size="small" variant="contained" startIcon={<Download />}>Download</Button>
        </CardActions>
      </Card>
    );
  }
  return (
    <Container sx={{
        marginTop: {xs:'5rem',md:'0rem'}
    }}>
      <Grid container gap={2} spacing={ {xl: 2}}>
        {Cards}
      </Grid>
    </Container>
  );
}
