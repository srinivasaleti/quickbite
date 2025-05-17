import React from "react";
import { Grid } from "./Styled";
import { Product } from "../Product/Product";
import { useProducts } from "../../hooks";

export const ProductList: React.FC = () => {
  const { products } = useProducts();

  return (
    <div>
      <h2>Desserts</h2>
      <Grid>
        {products?.map((product) => (
          <Product key={product.id} product={product} />
        ))}
      </Grid>
    </div>
  );
};
