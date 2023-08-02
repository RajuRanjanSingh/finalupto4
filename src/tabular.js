import React, { useState, useEffect } from 'react';
import axios from 'axios';
import './User.css';

function TabularTable() {
  const [users, setUsers] = useState([]);
  useEffect(() => {
    fetchUsers();
  }, []);

  const fetchUsers = async () => {
    try {
      const response = await axios.get('http://localhost:4747/getView');
      setUsers(response.data.data);
    } catch (error) {
      console.error('Error fetching users:', error);
    }
  };
  return (
    
    <>
   
    <div className='tabular-container'> 
      
      <table>
        <thead>
          <tr>
            <th>Name</th>
            <th>Total_Leave</th>
          </tr>
        </thead>
        <tbody>
          {users?.map((user) => (
            <tr key={user.id}>
              <td>{user.name}</td>
              <td>{user.total_leave}</td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
    <h2 className='tabular-head' >Tabular visualization</h2>
      
    </>
  );
}

export default TabularTable;
