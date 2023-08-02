import React from "react";
import "./App.css";
import LeaveForm from './User';
import Layout from "./Layout";
import { BrowserRouter, Routes, Route } from "react-router-dom";

import Tabulargraph from "./tabulargraph";
import Bargraph from "./bargraph";
import PieChart from "./PieChart";
import AllComponents1 from "./allcomponents";
function App() {
  return (
    <div className="App">
      <main>
        <BrowserRouter>
          <Routes>
            <Route path="/" element={<Layout />}>
              <Route index element={<LeaveForm />} />
              
              <Route path="tabular" element={<div className="tabulargraph"><Tabulargraph /></div>}/>
              <Route path="bar" element={<div className="bargraph"><Bargraph /></div>}/>
              <Route path="piegraph" element={<div className="piegraph"><PieChart /></div>}/>
              
              <Route path="Dashboard" element={<AllComponents1 />} />
            </Route>
          </Routes>
        </BrowserRouter>
      </main>
    </div>
  );
}

export default App;

