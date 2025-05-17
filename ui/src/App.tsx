import { ToastContainer } from "react-toastify";
import { Home } from "./pages/home";
import { ModalProvider } from "./common/components/Modal";

function App() {
  return (
    <>
      <ModalProvider>
        <Home />
        <ToastContainer />
      </ModalProvider>
    </>
  );
}

export default App;
