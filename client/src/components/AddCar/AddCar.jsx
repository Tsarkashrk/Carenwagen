import axios from 'axios'
import { useState } from 'react'

const AddCar = () => {
  const [formData, setFormData] = useState({
    title: '',
    model: '',
    make: '',
    year: 0,
    color: '',
    price: 0,
    mileage: 0,
    description: '',
    transmission: '',
    fuel_type: '',
    image_url: '',
  })

  const handleChange = (e) => {
    const { name, value } = e.target
    setFormData((prevData) => ({
      ...prevData,
      [name]: value,
    }))
  }

  const handleSubmit = async (e) => {
    e.preventDefault()

    try {
      const response = await axios.post('http://localhost:4000/v1/cars', formData)
      console.log(response.data)
    } catch (error) {
      console.error(error)
    }
  }

  return (
    <div>
      <h1>Create Car</h1>
      <form onSubmit={handleSubmit}>
        <div>
          <label>Title:</label>
          <input type="text" name="title" value={formData.title} onChange={handleChange} />
        </div>
        <div>
          <label>Model:</label>
          <input type="text" name="model" value={formData.model} onChange={handleChange} />
        </div>
        <div>
          <label>Make:</label>
          <input type="text" name="make" value={formData.make} onChange={handleChange} />
        </div>
        <div>
          <label>Year:</label>
          <input type="number" name="year" value={formData.year} onChange={handleChange} />
        </div>
        <div>
          <label>Color:</label>
          <input type="text" name="color" value={formData.color} onChange={handleChange} />
        </div>
        <div>
          <label>Price:</label>
          <input type="number" name="price" value={formData.price} onChange={handleChange} />
        </div>
        <div>
          <label>Mileage:</label>
          <input type="number" name="mileage" value={formData.mileage} onChange={handleChange} />
        </div>
        <div>
          <label>Description:</label>
          <input
            type="text"
            name="description"
            value={formData.description}
            onChange={handleChange}
          />
        </div>
        <div>
          <label>Transmission:</label>
          <input
            type="text"
            name="transmission"
            value={formData.transmission}
            onChange={handleChange}
          />
        </div>
        <div>
          <label>Fuel type:</label>
          <input type="text" name="fuel_type" value={formData.fuel_type} onChange={handleChange} />
        </div>
        <div>
          <label>Image url:</label>
          <input type="text" name="image_url" value={formData.image_url} onChange={handleChange} />
        </div>
        <button type="submit">Create Car</button>
      </form>
    </div>
  )
}

export default AddCar