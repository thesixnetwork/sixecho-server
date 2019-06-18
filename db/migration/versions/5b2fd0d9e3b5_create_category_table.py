"""create category table

Revision ID: 5b2fd0d9e3b5
Revises: ed0bc7456050
Create Date: 2019-06-10 18:09:40.540934

"""
import sqlalchemy as sa
from alembic import op

# revision identifiers, used by Alembic.
revision = '5b2fd0d9e3b5'
down_revision = 'ed0bc7456050'
branch_labels = None
depends_on = None


def upgrade():
    op.create_table(
        'categories',
        sa.Column('id', sa.Integer, primary_key=True),
        sa.Column('name', sa.String(255), nullable=False),
    )


def downgrade():
    op.drop_table('categories')
