# store

The _store_ implements a wrapper for the BoltDB and allows CRUD operations on generic items.

The _store_ supports configuring with _viper_ for work on _product / demo / etc_ base.

    viper.Set("BoltDBPath", "demo.db")
