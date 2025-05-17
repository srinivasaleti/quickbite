import React from "react";
import { Grid, Icon, NoProductsContainer } from "./Styled";
import { Product } from "../Product/Product";
import { useProducts } from "../../hooks";
import { Loader } from "../../../common/components/Loader";
import { Text } from "../../../common";

export const ProductList: React.FC = () => {
  const { loading, products } = useProducts();

  if (loading) {
    return (
      <div style={{ height: "100vh" }}>
        <Loader />
      </div>
    );
  }

  if (!products || !products.length) {
    return (
      <NoProductsContainer>
        <Icon>🍰</Icon>
        <Text> No products available</Text>
      </NoProductsContainer>
    );
  }

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
