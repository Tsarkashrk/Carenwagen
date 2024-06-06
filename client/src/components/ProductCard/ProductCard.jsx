/* eslint-disable react/prop-types */
import LinkButton from '../LinkButton/LinkButton'

const ProductCard = ({ carData }) => {
  return (
    <div className="product-card">
      <div className="product-card__wrapper">
        <div className="product-card__image-container">
          <img key={carData.id} src={`${carData.image}`} alt="" className="product-card__image" />
        </div>
        <div className="product-card__text">
          <div className="product-card__text--main">
            <h3 className="product-card__title">{carData.title}</h3>
            <p className="product-card__year">{carData.year}</p>
          </div>
          <p className="product-card__price">${carData.price}</p>
        </div>
        <LinkButton text="Details" link={`/car/${carData.id}/details`}/>
      </div>
    </div>
  )
}

export default ProductCard
