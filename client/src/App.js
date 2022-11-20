import Appbar from "./components/Navbar/Appbar";
import Login from "./components/Login/login";
import {BrowserRouter as Router, Routes, Route} from "react-router-dom";
import Signup from "./components/Signup/Signup";
import Home from "./components/Home/Home";
import Footer from "./components/Footer/Footer";
function App() {
  return (
    <>
    <Appbar></Appbar>
    <Router>
      <Routes>
        <Route path="/" element={<Home/>}></Route>
        <Route path="/login" element={<Login/>}></Route>
        <Route path="/signup" element={<Signup/>}></Route>
      </Routes>
    <Footer></Footer>
    </Router>
    </>
  );
}

export default App;
