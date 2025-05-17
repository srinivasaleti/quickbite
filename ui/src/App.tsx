import { ToastContainer } from "react-toastify";
import { Home } from "./pages/home";
import { ModalProvider } from "./common/components/Modal";
import { CartProvider } from "./cart/CartContext";

function App() {
  return (
    <>
      <ModalProvider>
        <CartProvider>
          <Home />
          <ToastContainer />
        </CartProvider>
      </ModalProvider>
    </>
  );
}

export default App;
