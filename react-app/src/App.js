import logo from './logo.svg';
import './App.css';
import { useQuery, gql } from '@apollo/client';

const LIST_SHOP = gql`
  query {
    listShop {
      id
      address
      phoneNumber
      notes
    }
  }
`;

function DisplayLocations() {
  const { loading, error, data } = useQuery(LIST_SHOP);

  if (loading) return <p>Loading...</p>;
  if (error) return <p>Error : {error.message}</p>;

  return data.listShop.map(({ id, address, phoneNumber, notes }) => (
    <div key={id}>
      <h3>Shop {id}</h3>
      <b>Address:</b>
      <p>{address}</p>
      <br />
      <b>Phone Number:</b>
      <p>{phoneNumber}</p>
      <br />
      <b>Notes:</b>
      <p>{notes}</p>
      <br />
    </div>
  ));
}

function App() {
  return (
    <div>
      <h2>My first Apollo app 🚀</h2>
      <br/>
      <DisplayLocations />
    </div>
  );
}

export default App;
