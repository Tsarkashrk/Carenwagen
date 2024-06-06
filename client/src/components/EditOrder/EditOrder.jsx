import React, { useState } from 'react'
import axios from 'axios'

const EditOrder = () => {
  const [orderId, setOrderId] = useState('')
  const [carId, setCarId] = useState('')
  const [message, setMessage] = useState('')

  const handleSubmit = async (e) => {
    e.preventDefault()

    const updatedOrder = {
      car_id: parseInt(carId),
    }

    try {
      const response = await axios.put(`http://localhost:4001/v1/order/${orderId}`, updatedOrder)
      setMessage('Order updated successfully!')
      console.log(response.data)
    } catch (error) {
      setMessage('Error updating order')
      console.error(error)
    }
  }

  return (
    <div>
      <h1>Edit Order</h1>
      <form onSubmit={handleSubmit}>
        <div>
          <label>Order Id:</label>
          <input type="number" value={orderId} onChange={(e) => setOrderId(e.target.value)} />
        </div>
        <div>
          <label>New Car Id:</label>
          <input type="number" value={carId} onChange={(e) => setCarId(e.target.value)} />
        </div>
        <button type="submit">Update Order</button>
      </form>
      {message && <p>{message}</p>}
    </div>
  )
}

export default EditOrder
