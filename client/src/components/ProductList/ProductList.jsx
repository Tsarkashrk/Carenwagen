/* eslint-disable react/jsx-key */
import ProductCard from '../ProductCard/ProductCard'

/* eslint-disable react/prop-types */
const ProductList = ({ carsData }) => {
  return (
    <div className="product-list">
      <div className="product-list__wrapper">
        {carsData.map((carData) => (
          <ProductCard key={carData.id} carData={carData} />
        ))}
      </div>
    </div>
  )
}

export default ProductList
