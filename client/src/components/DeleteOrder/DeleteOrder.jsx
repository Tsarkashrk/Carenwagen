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
      } catch (err) {
        setError(err.message)
      }
    }

    fetchOrders()
  }, [])

  const deleteOrder = async (id) => {
    try {
      await axios.delete(`http://localhost:4001/v1/order/${id}`)
      setOrders(orders.filter((order) => order.id !== id))
    } catch (err) {
      setError(err.message)
    }
  }

  if (error) {
    return <div>Error: {error}</div>
  }

  return (
    <div className="orders">
      <div className="orders__wrapper">
        <ul className="orders__list">
          {orders.map((order) => (
            <li key={order.id}>
              Order Id: {order.id}
              <button onClick={() => deleteOrder(order.id)}>Delete</button>
            </li>
          ))}
        </ul>
      </div>
    </div>
  )
}

export default Orders
