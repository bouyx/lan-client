import fetch from 'isomorphic-unfetch'

const Index = props => (
  <div>
    <div>
      <button>create</button>
    </div>
    <div>
      <button>details</button>
    </div>
    <div>
      <button onClick={props.prout}>prout</button>
      <div>prout ={props.pets}</div>


    </div>
  </div>
)
Index.prout = async function () {
  console.log("a kadok")
  const res = await fetch(`https://lanmanager.antoinebouyx.now.sh/go/prout`)
  const pets = await res.json()
  return { pets }
}

export default Index