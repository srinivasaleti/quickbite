import { createContext, useContext, useState, type ReactNode } from "react";
import styled from "styled-components";

const Overlay = styled.div`
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background-color: rgba(0, 0, 0, 0.5);
  z-index: 100;
`;

const ModalContainer = styled.div`
  position: fixed;
  bottom: 0;
  left: 0;
  right: 0;
  height: 70vh;
  background: white;
  border-top-left-radius: 15px;
  border-top-right-radius: 15px;
  z-index: 101;
  padding: 20px;
  box-shadow: 0 -2px 10px rgba(0, 0, 0, 0.3);
  overflow-y: auto;

  @media (min-width: 768px) {
    top: 50%;
    left: 50%;
    bottom: auto;
    right: auto;
    height: auto;
    max-height: 80vh;
    width: auto;
    min-width: 500px;
    transform: translate(-50%, -50%);
    border-radius: 15px;
  }
`;

const CloseIcon = styled.span`
  position: absolute;
  top: 10px;
  right: 15px;
  font-size: 24px;
  font-weight: bold;
  cursor: pointer;
`;

export default function Modal({ children }: { children: ReactNode }) {
  const { isOpen, closeModal } = useModal();

  if (!isOpen) return null;

  return (
    <>
      <Overlay onClick={closeModal} />
      <ModalContainer>
        <CloseIcon onClick={closeModal}>&times;</CloseIcon>
        {children}
      </ModalContainer>
    </>
  );
}

interface ModalContextType {
  isOpen: boolean;
  openModal: () => void;
  closeModal: () => void;
}

const ModalContext = createContext<ModalContextType | undefined>(undefined);

export function ModalProvider({ children }: { children: ReactNode }) {
  const [isOpen, setIsOpen] = useState(false);

  const openModal = () => setIsOpen(true);

  const closeModal = () => setIsOpen(false);

  return (
    <ModalContext.Provider value={{ isOpen, openModal, closeModal }}>
      {children}
    </ModalContext.Provider>
  );
}

export function useModal(): ModalContextType {
  const context = useContext(ModalContext);
  if (!context) {
    throw new Error("useModal must be used within a ModalProvider");
  }
  return context;
}
