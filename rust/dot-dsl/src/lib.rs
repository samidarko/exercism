pub mod graph {
    use crate::graph::graph_items::edge::Edge;
    use crate::graph::graph_items::node::Node;
    use std::collections::HashMap;

    // pub trait Attr {
    //     fn get_attrs(self) -> HashMap<String, String>;
    //     fn attr(&self, name: &str) -> Option<&str> {
    //         self.get_attrs().get(name).map(String::as_str)
    //     }
    // }

    pub mod graph_items {

        pub mod edge {
            use std::collections::HashMap;

            #[derive(Clone, PartialEq, Debug, Default)]
            pub struct Edge<'a> {
                start: &'a str,
                end: &'a str,
                attrs: HashMap<String, String>,
            }
            impl<'a> Edge<'a> {
                pub fn new(start: &'a str, end: &'a str) -> Self {
                    Self {
                        start,
                        end,
                        ..Default::default()
                    }
                }
                pub fn with_attrs(self, attrs: &[(&'a str, &'a str)]) -> Self {
                    let attrs = HashMap::from_iter(
                        attrs
                            .into_iter()
                            .map(|(name, value)| (name.to_string(), value.to_string())),
                    );
                    Self {
                        attrs,
                        ..self
                    }
                }

                pub fn attr(&self, name: &'a str) -> Option<&str> {
                    self.attrs.get(name).map(String::as_str)
                }
            }
        }
        pub mod node {
            use std::collections::HashMap;

            #[derive(Clone, PartialEq, Debug, Default)]
            pub struct Node<'a> {
                pub(crate) name: &'a str,
                pub(crate) attrs: HashMap<String, String>,
            }
            impl<'a> Node<'a> {
                pub fn new(name: &'a str) -> Self {
                    Self {
                        name,
                        ..Default::default()
                    }
                }
                pub fn with_attrs(self, attrs: &[(&'a str, &'a str)]) -> Self {
                    let attrs = HashMap::from_iter(
                        attrs
                            .into_iter()
                            .map(|(name, value)| (name.to_string(), value.to_string())),
                    );
                    Self {
                        attrs,
                        ..self
                    }
                }

                pub fn attr(&self, name: &'a str) -> Option<&str> {
                    self.attrs.get(name).map(String::as_str)
                }
            }
        }
        // impl Attr for Node {
        //     // fn get_attrs<'a>(&self) -> Vec<(&'a str, &'a str)>;
        //     fn get_attrs(self) -> Vec<(&'a str, &'a str)> {
        //         self.attrs
        //     }
        // }
    }
    #[derive(Default)]
    pub struct Graph<'a> {
        pub nodes: Vec<Node<'a>>,
        pub edges: Vec<Edge<'a>>,
        pub attrs: HashMap<String, String>,
    }

    impl<'a> Graph<'a> {
        pub fn new() -> Self {
            Self::default()
        }

        pub fn with_nodes(self, nodes: &[Node<'a>]) -> Self {
            Self {
                nodes: nodes.to_vec(),
                ..self
            }
        }

        pub fn with_edges(self, edges: &[Edge<'a>]) -> Self {
            Self {
                edges: edges.to_vec(),
                ..self
            }
        }

        pub fn with_attrs(self, attrs: &[(&'a str, &'a str)]) -> Self {
            let attrs = HashMap::from_iter(
                attrs
                    .into_iter()
                    .map(|(name, value)| (name.to_string(), value.to_string())),
            );
            Self { attrs, ..self }
        }

        pub fn node(&self, name: &str) -> Option<&Node> {
            self.nodes.iter().find(|node| node.name == name)
        }
    }
}
