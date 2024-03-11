import React, {useReducer} from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import reportWebVitals from './reportWebVitals';
import App from "./App";
import {reducer, initialState} from "./store/reducer";

const root = ReactDOM.createRoot(
  document.getElementById('root') as HTMLElement
);

function AppReducer(props: any) {
  const [state, dispatch] = useReducer(reducer, initialState);

  return <App {...props} {...{state, dispatch}} />;
}

root.render(
  <React.StrictMode>
    <AppReducer/>
  </React.StrictMode>
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
