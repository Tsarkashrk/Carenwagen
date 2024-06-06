import Input from '../Input/Input'
import inputTitles from '../../constants/inputTitles'

const AddCar = () => {
  return (
    <div className="add-car">
      <div className="add-car__wrapper">
        {inputTitles.map((inputTitle) => (
          <div key={inputTitle.id} className="add-car__item">
            <h3 className="add-car__title">{inputTitle.title}</h3>
            <Input text={inputTitle.title} />
          </div>
        ))}
      </div>
    </div>
  )
}

export default AddCar
