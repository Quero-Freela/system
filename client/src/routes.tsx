import React from 'react';
import {BrowserRouter} from "react-router-dom";
import CasesRoutes from './cases';

export default function routes() {
  return (
    <BrowserRouter>
      <CasesRoutes />
    </BrowserRouter>
  );
}
