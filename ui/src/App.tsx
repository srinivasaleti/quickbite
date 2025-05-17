import { ToastContainer } from "react-toastify";
import { ProductList } from "./product/components/ProductList";

function App() {
  return (
    <>
      <ProductList />
      <ToastContainer />
    </>
  );
}

export default App;
