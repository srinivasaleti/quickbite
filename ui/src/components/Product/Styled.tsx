import styled from "styled-components";

export const Card = styled.div`
  width: 200px;
  border: 1px solid #ddd;
  border-radius: 12px;
  padding: 12px;
  box-shadow: 0 2px 5px rgba(0, 0, 0, 0.1);
`;

export const Image = styled.img`
  width: 100%;
  height: 180px;
  object-fit: cover;
  border-radius: 8px;
`;

export const Name = styled.h3`
  margin: 8px 0 4px;
  font-size: 16px;
`;

export const Category = styled.p`
  color: gray;
  font-size: 13px;
`;

export const Price = styled.p`
  font-weight: bold;
  margin-top: 4px;
`;
