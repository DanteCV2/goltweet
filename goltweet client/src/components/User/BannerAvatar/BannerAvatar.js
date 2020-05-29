import React, { useState, useEffect } from "react";
import { Button } from "react-bootstrap";
import ConfigModal from "../../Modal/ConfigModal";
import EditUserForm from "../EditUserForm";
import AvatarNotFound from "../../../assets/png/avatar.png";
import { API_HOST } from "../../../utils/constant";
import {
  checkFollowApi,
  folllowUserApi,
  unfolllowUserApi,
} from "../../../api/follow";

import "./BannerAvatar.scss";

export default function BannerAvatar(props) {
  const { user, loggedUser } = props;
  const [showModal, setShowModal] = useState(false);
  const [following, setFollowing] = useState(null);
  const [reloadFollow, setReloadFollow] = useState(false);
  const bannerUrl = user?.banner ? `${API_HOST}/getBanner?id=${user.id}` : null;
  const avatarUrl = user?.avatar
    ? `${API_HOST}/getAvatar?id=${user.id}`
    : AvatarNotFound;

  useEffect(() => {
    if (user) {
      checkFollowApi(user?.id).then((response) => {
        if (response?.status) {
          setFollowing(true);
        } else {
          setFollowing(false);
        }
      });
    }
    setReloadFollow(false);
  }, [user, reloadFollow]);

  const onFollow = () => {
    folllowUserApi(user.id).then(() => {
      setReloadFollow(true);
    });
  };

  const onUnfollow = () => {
    unfolllowUserApi(user.id).then(() => {
      setReloadFollow(true);
    });
  };

  return (
    <div
      className="banner-avatar"
      style={{ backgroundImage: `url('${bannerUrl}')` }}
    >
      <div
        className="avatar"
        style={{ backgroundImage: `url('${avatarUrl}')` }}
      />
      {user && (
        <div className="options">
          {loggedUser._id === user.id && (
            <Button onClick={() => setShowModal(true)}>Editar perfil</Button>
          )}

          {loggedUser._id !== user.id &&
            following !== null &&
            (following ? (
              <Button className="unfollow" onClick={onUnfollow}>
                <span>Siguiendo</span>
              </Button>
            ) : (
              <Button onClick={onFollow}>Seguir</Button>
            ))}
        </div>
      )}
      <ConfigModal
        show={showModal}
        setShow={setShowModal}
        title="Editar Perfil"
      >
        <EditUserForm user={user} setShowModal={setShowModal} />
      </ConfigModal>
    </div>
  );
}