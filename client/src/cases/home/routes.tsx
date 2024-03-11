import React from "react";
import { Route } from "react-router-dom";
import { Home } from "./pages";

export default function routes() {
  return [
    <Route index path="/" element={<Home/>} key="/"/>
  ];
}
