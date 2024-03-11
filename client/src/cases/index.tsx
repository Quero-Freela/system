import React from 'react';
import {HomeRoutes} from './home';
import {ProjectsRoutes} from './projects';
import {SecurityRoutes} from './security';
import {Routes} from "react-router-dom";

const CasesRoutes = [
  ...HomeRoutes(),
  ...ProjectsRoutes(),
  ...SecurityRoutes()
];

export default function AllRoutes() {
  return (
    <Routes>
      {CasesRoutes}
    </Routes>
  );
}
