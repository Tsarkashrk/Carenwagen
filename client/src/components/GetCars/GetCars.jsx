import React, { useState, useEffect } from 'react';
import axios from 'axios';

const GetCars = () => {
  const [cars, setCars] = useState([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const fetchCars = async () => {
      try {
        const response = await axios.get('http://localhost:4000/v1/cars');
        setCars(response.data.cars);
        console.log(cars)
        setLoading(false);
      } catch (error) {
        console.error('Error fetching cars:', error);
        setLoading(false);
      }
    };

    fetchCars();
  }, []);

  if (loading) {
    return <div>Loading...</div>;
  }

  return (
    <div>
      <h1>Cars List</h1>
      <ul>
        {cars.map((car) => (
          <li key={car.id}>
            <p>Title: {car.title}</p>
            <p>Model: {car.model}</p>
            <p>Make: {car.make}</p>
            <p>Year: {car.year}</p>
          </li>
        ))}
      </ul>
    </div>
  );
};

export default GetCars;
