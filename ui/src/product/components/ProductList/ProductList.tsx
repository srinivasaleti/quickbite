import React from "react";
import { Grid, Icon, NoProductsContainer } from "./Styled";
import { Product } from "../Product/Product";
import { useProducts } from "../../hooks";
import { Loader } from "../../../common/components/Loader";
import { Text } from "../../../common";
import { useTheme } from "styled-components";

export const ProductList: React.FC = () => {
  const { loading, products } = useProducts();
  const { fontWeights } = useTheme();
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
      <Text size="36px" weight={fontWeights.semiBold}>
        Desserts
      </Text>
      <Grid>
        {products?.map((product) => (
          <Product key={product.id} product={product} />
        ))}
      </Grid>
    </div>
  );
};
