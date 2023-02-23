import React from 'react';
import { Container, Button } from 'react-bootstrap';
import {useDispatch, useSelector} from "react-redux";
import { FetchRankings } from '../Controllers/Auth/user';

const RankingsPage = () => {
    const dispatch = useDispatch();
  const rankings = useSelector(state => state.client.ranking);

  return (
    <Container className="p-0">
      <div className="container mt-5">
        <h1 className="text-center mb-5">Rankings</h1>
        <Button onClick={() => dispatch(FetchRankings())}>Refresh Rankings</Button>
        <table className="table table-striped">
          <thead>
            <tr>
              <th>Rank</th>
              <th>Name</th>
              <th>Points</th>
              <th>Language</th>
            </tr>
          </thead>
          <tbody>
            {rankings.map(({ rank, name, points, language }) => (
              <tr key={rank}>
                <td>{rank}</td>
                <td>{name}</td>
                <td>{points}</td>
                <td>{language}</td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </Container>
  );
};

export default RankingsPage;