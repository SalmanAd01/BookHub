import * as React from "react";
import AppBar from "@mui/material/AppBar";
import Box from "@mui/material/Box";
import Toolbar from "@mui/material/Toolbar";
import Typography from "@mui/material/Typography";
import Button from "@mui/material/Button";
import IconButton from "@mui/material/IconButton";
import MenuIcon from "@mui/icons-material/Menu";
import Drawer from "@mui/material/Drawer";
import Sidebar from "./Sidebar";
export default function Appbar() {
  const [drawerOpen, setDrawerOpen] = React.useState(false);
  return (
    <Box>
      <AppBar position="static">
        <Toolbar>
          <IconButton
            size="large"
            edge="start"
            color="inherit"
            aria-label="menu"
            sx={{ mr: 2 }}
            onClick={() => setDrawerOpen(true)}
          >
            <MenuIcon  />
          </IconButton>

          <Typography variant="h6" component="div" sx={{ flexGrow: 1 }}>
            Bookhub
          </Typography>
          <Button color="inherit" href="/login">Login</Button>
        </Toolbar>
        <Drawer anchor={"left"} open={drawerOpen} onClose={() => setDrawerOpen(false)}>
          <Sidebar></Sidebar>
        </Drawer>
      </AppBar>
    </Box>
  );
}
