import React, { useEffect, useState } from 'react';
import axios from 'axios';

const Orders = () => {
  const [orders, setOrders] = useState([]);
  const [error, setError] = useState('');

  useEffect(() => {
    const fetchOrders = async () => {
      try {
        const response = await axios.get('http://localhost:4001/v1/order');
        setOrders(response.data.orders);
        console.log(orders)
      } catch (err) {
        setError(err.message);
      }
    };

    fetchOrders();
  }, []);

  if (error) {
    return <div>Error: {error}</div>;
  }

  return (
    <div>
      <h1>Orders</h1>
      <ul>
        {orders.map(order => (
          <li key={order.id}>{order.}</li>
        ))}
      </ul>
    </div>
  );
};

export default Orders;
