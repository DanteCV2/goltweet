import React from "react";
import { Modal } from "react-bootstrap";
import LogoWhite from "../../../assets/png/original-2.png";
import "./BasicModal.scss";

export default function BasicModal(props) {
  const { show, setShowModal, children } = props;
  return (
    <Modal
      className="basic-modal"
      show={show}
      onHide={() => {
        setShowModal(false);
      }}
      centered
      sieze="lg"
    >
      <Modal.Header>
        <Modal.Title>
          <img src={LogoWhite} alt="goTweet" />
        </Modal.Title>
      </Modal.Header>
      <Modal.Body>{children}</Modal.Body>
    </Modal>
  );
}
