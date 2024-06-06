import ProductList from '../../components/ProductList/ProductList'
import Heading from '../../components/Heading/Heading'
import CARS_DATA from '../../constants/carsData'

const Catalog = () => {
  return (
    <div className="catalog">
      <div className="catalog__wrapper">
        <Heading text="catalog" />
        <ProductList carsData={CARS_DATA} />
      </div>
    </div>
  )
}

export default Catalog
