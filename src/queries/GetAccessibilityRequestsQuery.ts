import { gql } from '@apollo/client';

export default gql`
  query GetAccessibilityRequests($first: Int!) {
    accessibilityRequests(first: $first) {
      edges {
        node {
          id
          submittedAt
          system {
            name
            lcid
            businessOwner {
              name
              component
            }
          }
        }
      }
    }
  }
`;
