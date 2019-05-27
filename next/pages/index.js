import React from 'react'
import fetch from 'isomorphic-unfetch'

const fetchPets = async () => {
  console.log('a kadok')
  const res = await fetch(`https://lanmanager.antoinebouyx.now.sh/go/prout`)
  const pets = await res.json()
  return { pets }
}

function Index() {
  const [pets, setPets] = React.useState()
  const handleOnClick = async () => {
    const _pets = await fetchPets()
    setPets(_pets)
  }
  return (
    <div>
      <div>
        <button>create</button>
      </div>
      <div>
        <button>details</button>
      </div>
      <div>
        <button onClick={handleOnClick}>prout</button>
        <pre>{JSON.stringify(pets, null, 2)}</pre>
      </div>
    </div>
  )
}

export default Index