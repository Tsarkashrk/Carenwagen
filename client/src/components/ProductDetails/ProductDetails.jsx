import { useParams } from 'react-router-dom';
import CARS_DATA from '../../constants/carsData';

const ProductDetails = () => {
  const { id } = useParams();

  const car = CARS_DATA.find((car) => car.id === parseInt(id));

  return (
    <div className="product-details">
      <div className="product-details__wrapper">
        {car ? (
          <div>
            <h2>{car.title}</h2>
            <p>Year: {car.year}</p>
            <p>Price: ${car.price}</p>
            <img src={car.image} alt={car.title} />
          </div>
        ) : (
          <p>Car not found</p>
        )}
      </div>
    </div>
  );
};

export default ProductDetails;
