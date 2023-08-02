import React, { useState, useEffect } from 'react';
import axios from 'axios';
import { BarChart, Bar, XAxis, YAxis, CartesianGrid, Tooltip, Legend } from 'recharts';
const BarGraph = () => {
  const [users, setUsers] = useState([]);
  useEffect(() => {
    fetchUsers();
  }, []);
  const fetchUsers = async () => {
    try {
      const response = await axios.get('http://localhost:4747/getkpi4');
      setUsers(response.data.data);
    } catch (error) {
      console.error('Error fetching users:', error);
    }
  };
  // Extract data for the bar graph
  const chartData = users.map((user) => ({
    name: user.manager,
    total_leave: user.total_leave,
  }));
  return (
    <div className='divide'>
      <BarChart width={600} height={400} data={chartData} margin={{ top: 20, right: 30, left: 20, bottom: 5 }}>
        <CartesianGrid strokeDasharray='3 3' />
        <XAxis dataKey='name' />
        <YAxis />
        <Tooltip />
        <Legend />
        <Bar dataKey='total_leave' fill='rgba(75, 192, 192, 0.6)' />
      </BarChart>
      <h2 className='bar-head'>Bar chart</h2>
    </div>
  );
};
export default BarGraph;