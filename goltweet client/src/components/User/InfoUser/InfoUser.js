import React from "react";
import moment from "moment";
import localization from "moment/locale/es";
import { Location, Link, Calendar } from "../../../utils/icons";

import "./InfoUser.scss";

export default function InfoUser(props) {
  const { user } = props;
  return (
    <div className="info-user">
      <h2 className="name">
        {user?.name} {user?.surname}
      </h2>
      <p className="email">{user?.email}</p>
      {user?.biography && <div className="biography">{user.biography}</div>}
      <div className="more-info">
        {user?.location && (
          <p>
            <Location />
            {user.location}
          </p>
        )}
        {user?.website && (
          <a
            href={user.website}
            alt={user.website}
            target="_blank"
            rel="noopener noreferrer"
          >
            <Link /> {user.website}
          </a>
        )}
        {user?.birthDate && (
          <p>
            <Calendar />
            {moment(user.birthDate).locale("es", localization).format("LL")}
          </p>
        )}
      </div>
    </div>
  );
}
