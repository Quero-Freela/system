import React from "react";
import { Route } from "react-router-dom";
import { Home, Login } from "./pages";

export default function routes() {
  return [
    <Route path="/security" element={<Home/>} key="/security">
      <Route path="login" element={<Login/>} key="/security/login" />
    </Route>
  ];
}
