import React, { useEffect, useState } from "react";
import { PieChart, Pie, Cell, Tooltip, Legend } from "recharts";
const PieChartcomponent = () => {
    const [data, setData] = useState([]);
    useEffect(() => {
        fetchData();
    }, []);
    const fetchData = async () => {
        try {
            const response = await fetch("http://localhost:4747/getkpi6");
            const jsonData = await response.json();
            setData(jsonData.data);
        } catch (error) {
            console.error("Error fetching data:", error);
        }
    };
    const COLORS = ["#0088FE", "#00C49F", "#FFBB28", "#FF8042", "#8884d8"];

    return (
        <>
     
       <div className="piechart1">
        <PieChart width={400} height={370} >
            <Pie     
                data={data.filter((item) => item.Team_Name === "AI")}
                dataKey="Leave_count"
                nameKey="Leave_type"
                cx="50%"
                cy="50%"
                outerRadius={80}
                fill="#8884d8"
                label={(entry) => entry.Leave_type}
            >
                {data.map((entry, index) => (
                    <Cell key={`cell-${index}`} fill={COLORS[index % COLORS.length]} />
                ))}
            </Pie>
            <Tooltip />
            <Legend />

        </PieChart>
        <h2 className="AI">Team AI graph</h2>
        </div>
      
        <div className="piechart2"> 
        <PieChart width={400} height={370} >
            <Pie   
                data={data.filter((item) => item.Team_Name === "IT")}
                dataKey="Leave_count"
                nameKey="Leave_type"
                cx="50%"
                cy="50%"
                outerRadius={80}
                fill="#8884d8"
                label={(entry) => entry.Leave_type}
            >
                {data.map((entry, index) => (
                    <Cell key={`cell-${index}`} fill={COLORS[index % COLORS.length]} />
                ))}
            </Pie>
            <Tooltip />
            <Legend />
        </PieChart>
        <h2 className="IT">Team IT graph</h2>
      </div>
        </>
    );
};
export default PieChartcomponent;