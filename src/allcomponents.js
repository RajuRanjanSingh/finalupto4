import PieChart from "./PieChart";
import BarChat from "./BarChat";
import './User.css';
import React from "react";
import TabularTable from "./tabular";




function AllComponents1 (){
    return(
     <>
        <div id="main" >
        <div> <PieChart/></div>
        <div>
            <TabularTable/>
            <BarChat/>
        </div>
       
        
        </div>
        
       
          </>
    );
}

export default AllComponents1;