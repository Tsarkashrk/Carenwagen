import React from 'react'
import Input from '../Input/Input'
import inputTitles from '../../constants/inputTitles'

const UpdateCar = () => {
  return (
    <div className="update-car">
      <div className="update-car__wrapper">
        {inputTitles.map((inputTitle) => (
          <div key={inputTitle.id} className="update-car__item">
            <h3 className="update-car__title">{inputTitle.title}</h3>
            <Input text={inputTitle.title} />
          </div>
        ))}
      </div>
    </div>
  )
}

export default UpdateCar
