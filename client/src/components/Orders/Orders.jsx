import React, { useEffect, useState } from 'react'
import axios from 'axios'

const Orders = () => {
  const [orders, setOrders] = useState([])
  const [error, setError] = useState('')

  useEffect(() => {
    const fetchOrders = async () => {
      try {
        const response = await axios.get('http://localhost:4001/v1/order')
        setOrders(response.data.orders)
        console.log(orders)
      } catch (err) {
        setError(err.message)
        console.log(err);
      }
    }

    fetchOrders()
  }, [])

  if (error) {
    return <div>Error: {error}</div>
  }

  return (
    <div className="orders">
      <div className="orders__wrapper">
        <ul className="orders__list">
          {orders.map((order) => (
            <li key={order.id} className="orders__item">
              <p className="orders__element">Object Number: {order.id}</p>
              <p className="orders__id">Order ID: {order.id}</p>
              <p className="orders__customer-id">Customer ID: {order.customer_id}</p>
              <p className="orders__car-id">Car ID: {order.car_id}</p>
            </li>
          ))}
        </ul>
      </div>
    </div>
  )
}

export default Orders
