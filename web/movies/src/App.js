import {  
  createBrowserRouter,
  RouterProvider,
  createRoutesFromElements,
  Route,
} from "react-router-dom";

import Home from "./components/Home";
import Movies from "./components/Movies";
import Movie from "./components/Movie";

//Layout
import RootLayout from "./components/RootLayout";
import ErrorPage from "./components/ErrorPage";
import Login from "./components/Login";
import ManageCatalog from "./components/ManageCatalog";
import EditMovie from "./components/EditMovie";

function App() {

  const router = createBrowserRouter(
    createRoutesFromElements(
      <Route path="/" element={<RootLayout/>} errorElement={<ErrorPage/>} >
        <Route index element={<Home/>}></Route>
        <Route path="movies" element={<Movies/>}></Route>
        <Route path="movie/:id" element={<Movie/>}></Route>
        <Route path="login" element={<Login/>}></Route>
        <Route path="admin" element={<ManageCatalog/>}></Route>
        <Route path="admin/movie" element={<EditMovie/>}></Route>
        <Route path="admin/movies/:id" element={<EditMovie/>}></Route>

      </Route>
    )
  );


  return (
    <RouterProvider router={router} />
  );
}

export default App;
