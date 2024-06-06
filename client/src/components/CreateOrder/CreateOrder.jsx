import React, { useState } from 'react'
import axios from 'axios'

const CreateOrder = () => {
  const [customerId, setCustomerId] = useState(0)
  const [carId, setCarId] = useState(0)

  const handleSubmit = async (e) => {
    e.preventDefault()

    const newOrder = {
      customer_id: customerId,
      car_id: carId,
    }

    try {
      const response = await axios.post('http://localhost:4001/v1/order', newOrder)
      console.log(response.data)
    } catch (error) {
      console.error(error)
    }
  }

  return (
    <div>
      <h1>Create Order</h1>
      <form onSubmit={handleSubmit}>
        <div>
          <label>Customer Id:</label>
          <input type="text" value={customerId} onChange={(e) => setCustomerId(e.target.value)} />
        </div>
        <div>
          <label>Car Id:</label>
          <input type="text" value={carId} onChange={(e) => setCarId(e.target.value)} />
        </div>
        <button type="submit">Create Order</button>
      </form>
    </div>
  )
}

export default CreateOrder
